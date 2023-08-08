import React, { useContext } from 'react';
import { AppContext } from '../context/AppContext'; 
const Remaining = () => {
    const { expenses, budget, currency } = useContext(AppContext);
    const totalExpenses = expenses.reduce((total, item) => {
        return (total = total + item.cost);
    }, 0);
    const alertType = totalExpenses > budget ? 'alert-danger' : 'alert-success';

    // 根據所選的貨幣格式化顯示
    const formatCurrency = (amount) => {
        return currency + amount;
    };

    return (
        <div className={`alert ${alertType}`}>
            <span>Remaining: {formatCurrency(budget - totalExpenses)}</span>
        </div>
    );
};
export default Remaining;
