// @ts-check
import { test, expect } from '@playwright/test';

test('First page load', async ({ page }) => {
  await page.goto('http://reactjs:80/');

  // Expect #root > div.card > button to contain text 'count is 0'
  await expect(page.locator('#root > div.card > button')).toHaveText('count is 0');
});


test('Click 2 times', async ({ page }) => {
  await page.goto('http://reactjs:80/');

  // Click #root > div.card > button twice
  await page.locator('#root > div.card > button').click();
  await page.locator('#root > div.card > button').click();

  // Expect #root > div.card > button to contain text 'count is 2'
  await expect(page.locator('#root > div.card > button')).toHaveText('count is 2');
});

