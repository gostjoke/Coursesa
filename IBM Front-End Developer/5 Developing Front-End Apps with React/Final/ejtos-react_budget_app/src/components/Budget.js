import React, { useContext } from 'react';
import { AppContext } from '../context/AppContext';

const Budget = () => {
    const { budget, dispatch, currency } = useContext(AppContext);

    const handleBudgetChange = (event) => {
        const newBudget = parseInt(event.target.value, 10);
        dispatch({ type: 'SET_BUDGET', payload: newBudget });
    };

    // 根據所選的貨幣格式化顯示
    const formatCurrency = (amount) => {
        return currency + amount;
    };

    return (
        <div className='alert alert-secondary'>
            <span>Budget: {formatCurrency(budget)}</span>
            <input
                required='required'
                type='number'
                id='budget'
                value={budget}
                style={{ marginLeft: '2rem' }}
                onChange={handleBudgetChange}
            />
        </div>
    );
};

export default Budget;
