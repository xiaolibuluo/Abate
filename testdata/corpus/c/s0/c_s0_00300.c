
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant208_0(int value, ...) {
    return value + 90;
}
#else
int variant208_0(int value) {
    return value + 90;
}
#endif
enum Kind208_1 { KIND_EMPTY208_1 = 0, KIND_FULL208_1 = 97 };

union Cell208_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind208_1(union Cell208_1 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL208_1 : KIND_EMPTY208_1;
}
struct Node208_2 {
    int value;
    struct Node208_2 *next;
};

int length208_2(const struct Node208_2 *head) {
    int count = 0;
    const struct Node208_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
