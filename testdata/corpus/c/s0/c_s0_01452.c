
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant987_0(int value, ...) {
    return value + 32;
}
#else
int variant987_0(int value) {
    return value + 32;
}
#endif
struct Node987_1 {
    int value;
    struct Node987_1 *next;
};

int length987_1(const struct Node987_1 *head) {
    int count = 0;
    const struct Node987_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
enum Kind987_2 { KIND_EMPTY987_2 = 0, KIND_FULL987_2 = 46 };

union Cell987_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind987_2(union Cell987_2 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL987_2 : KIND_EMPTY987_2;
}
