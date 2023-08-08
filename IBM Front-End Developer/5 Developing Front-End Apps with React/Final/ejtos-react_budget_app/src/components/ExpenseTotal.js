import React, { useContext } from 'react';
import { AppContext } from '../context/AppContext';
const ExpenseTotal = () => {
    const { expenses, currency } = useContext(AppContext);
    const totalExpenses = expenses.reduce((total, item) => {
        return (total += item.cost);
    }, 0);

    // 根據所選的貨幣格式化顯示
    const formatCurrency = (amount) => {
        return currency + amount;
    };

    return (
        <div className='alert alert-primary'>
            <span>Spent so far: {formatCurrency(totalExpenses)}</span>
        </div>
    );
};
export default ExpenseTotal;
