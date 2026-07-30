# Commerce AI Platform Evaluation Dataset

## Company Information

### 1. Who is the CEO?

Expected Tool:
- search_documents

Expected Response:
- CEO is Anup Mayank.

---

### 2. Who is the CTO?

Expected Tool:
- search_documents

Expected Response:
- CTO is Kundan Mayank Jha.

---

### 3. Who is the CFO?

Expected Tool:
- search_documents

Expected Response:
- CFO is Kumari Anshu.

---

### 4. Where is your headquarters?

Expected Tool:
- search_documents

Expected Response:
- Return headquarters information.

---

### 5. Tell me about your leadership team.

Expected Tool:
- search_documents

Expected Response:
- CEO
- CTO
- CFO
- COO

---

## Product Search

### 6. Do you sell blue t-shirts?

Expected Tool:
- search_products

Expected Response:
- List matching blue t-shirts.

---

### 7. Show me men's hoodies.

Expected Tool:
- search_products

Expected Response:
- Matching hoodies.

---

### 8. I'm looking for black jeans.

Expected Tool:
- search_products

Expected Response:
- Matching jeans.

---

### 9. Show waterproof jackets.

Expected Tool:
- search_products

Expected Response:
- Matching jackets.

---

### 10. Do you sell organic cotton clothing?

Expected Tool:
- search_products
- search_documents

Expected Response:
- Matching sustainable products.

---

## Inventory

### 11. Do you have NP-HZS-NVY-XL?

Expected Tool:
- check_inventory

Expected Response:
- SKU available
- Total quantity
- Warehouse quantities

---

### 12. Is SKU NP-TSH-BLU-M available?

Expected Tool:
- check_inventory

Expected Response:
- Inventory status.

---

### 13. Which warehouse has the most inventory for NP-HZS-NVY-XL?

Expected Tool:
- check_inventory

Expected Response:
- Los Angeles warehouse has highest quantity.

---

### 14. Is this product out of stock?

Expected Tool:
- check_inventory

Expected Response:
- Stock status.

---

### 15. How many units are available?

Expected Tool:
- check_inventory

Expected Response:
- Total available quantity.

---

## Returns

### 16. Can I return my order after 20 days?

Expected Tool:
- search_documents

Expected Response:
- Return policy.

---

### 17. Can I exchange a shirt?

Expected Tool:
- search_documents

Expected Response:
- Exchange policy.

---

### 18. How long does refund take?

Expected Tool:
- search_documents

Expected Response:
- Refund timeline.

---

### 19. Can I return clearance items?

Expected Tool:
- search_documents

Expected Response:
- Clearance policy.

---

### 20. What if I received a damaged item?

Expected Tool:
- search_documents

Expected Response:
- Damaged product process.

---

## Warranty

### 21. Does warranty cover broken zipper?

Expected Tool:
- search_documents

Expected Response:
- Covered.

---

### 22. Does warranty cover accidental damage?

Expected Tool:
- search_documents

Expected Response:
- Not covered.

---

### 23. How long is product warranty?

Expected Tool:
- search_documents

Expected Response:
- One year.

---

### 24. How do I file a warranty claim?

Expected Tool:
- search_documents

Expected Response:
- Claim procedure.

---

### 25. Is fading covered?

Expected Tool:
- search_documents

Expected Response:
- No.

---

## Shipping

### 26. Do you ship internationally?

Expected Tool:
- search_documents

Expected Response:
- International shipping policy.

---

### 27. How much is express shipping?

Expected Tool:
- search_documents

Expected Response:
- Express shipping details.

---

### 28. Do you offer same-day delivery?

Expected Tool:
- search_documents

Expected Response:
- Same-day delivery information.

---

### 29. How long does shipping take?

Expected Tool:
- search_documents

Expected Response:
- Shipping timeline.

---

### 30. Can I track my package?

Expected Tool:
- search_documents

Expected Response:
- Tracking process.

---

## Loyalty

### 31. How do I become Platinum?

Expected Tool:
- search_documents

Expected Response:
- Platinum requirements.

---

### 32. Do loyalty points expire?

Expected Tool:
- search_documents

Expected Response:
- Expiration rules.

---

### 33. Can I transfer points?

Expected Tool:
- search_documents

Expected Response:
- No.

---

### 34. What are Gold member benefits?

Expected Tool:
- search_documents

Expected Response:
- Gold benefits.

---

### 35. How do I redeem points?

Expected Tool:
- search_documents

Expected Response:
- Redemption process.

---

## Promotions

### 36. Do students get discounts?

Expected Tool:
- search_documents

Expected Response:
- Student discount.

---

### 37. Can I combine coupon codes?

Expected Tool:
- search_documents

Expected Response:
- Coupon policy.

---

### 38. Do military personnel get discounts?

Expected Tool:
- search_documents

Expected Response:
- Military discount.

---

### 39. Do you price match Amazon?

Expected Tool:
- search_documents

Expected Response:
- No price matching.

---

### 40. What promotions are running?

Expected Tool:
- search_documents

Expected Response:
- Seasonal promotions.

---

## Gift Cards

### 41. Do gift cards expire?

Expected Tool:
- search_documents

Expected Response:
- Never expire.

---

### 42. Can I reload a gift card?

Expected Tool:
- search_documents

Expected Response:
- No.

---

### 43. Can I use multiple gift cards?

Expected Tool:
- search_documents

Expected Response:
- Yes.

---

### 44. Can gift cards be refunded?

Expected Tool:
- search_documents

Expected Response:
- Non-refundable.

---

### 45. What gift card values are available?

Expected Tool:
- search_documents

Expected Response:
- $25, $50, $100, $250, $500.

---

## Multi-tool Scenarios

### 46. Do you have blue hoodies in stock?

Expected Tool:
- search_products
- check_inventory

Expected Response:
- Matching products with inventory.

---

### 47. Which blue t-shirt is available in Chicago?

Expected Tool:
- search_products
- check_inventory

Expected Response:
- Matching products available in Chicago warehouse.

---

### 48. Can I return a hoodie purchased with a student discount?

Expected Tool:
- search_documents

Expected Response:
- Return policy independent of promotion.

---

### 49. Does Platinum membership include free express shipping?

Expected Tool:
- search_documents

Expected Response:
- Platinum benefits include express shipping.

---

### 50. I have 300 loyalty points. Can I use them together with a gift card?

Expected Tool:
- search_documents

Expected Response:
- Rewards and gift cards can be used together if eligible.