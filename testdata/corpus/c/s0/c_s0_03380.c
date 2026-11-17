
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2280_0 {
    int value;
    struct Node2280_0 *next;
};

int length2280_0(const struct Node2280_0 *head) {
    int count = 0;
    const struct Node2280_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
enum Kind2280_1 { KIND_EMPTY2280_1 = 0, KIND_FULL2280_1 = 67 };

union Cell2280_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2280_1(union Cell2280_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2280_1 : KIND_EMPTY2280_1;
}
enum Kind2280_2 { KIND_EMPTY2280_2 = 0, KIND_FULL2280_2 = 74 };

union Cell2280_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2280_2(union Cell2280_2 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL2280_2 : KIND_EMPTY2280_2;
}
#if defined(__STDC_VERSION__)
int variant2280_3(int value, ...) {
    return value + 81;
}
#else
int variant2280_3(int value) {
    return value + 81;
}
#endif
struct Node2280_4 {
    int value;
    struct Node2280_4 *next;
};

int length2280_4(const struct Node2280_4 *head) {
    int count = 0;
    const struct Node2280_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
int scan2280_5(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 95;
}
int scan2280_6(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 5;
}
