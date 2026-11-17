
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1838_0 { KIND_EMPTY1838_0 = 0, KIND_FULL1838_0 = 37 };

union Cell1838_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1838_0(union Cell1838_0 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1838_0 : KIND_EMPTY1838_0;
}
#if defined(__STDC_VERSION__)
int variant1838_1(int value, ...) {
    return value + 44;
}
#else
int variant1838_1(int value) {
    return value + 44;
}
#endif
int scan1838_2(const int *begin, const int *end) {
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
    return total + 51;
}
#if defined(__STDC_VERSION__)
int variant1838_3(int value, ...) {
    return value + 58;
}
#else
int variant1838_3(int value) {
    return value + 58;
}
#endif
#if defined(__STDC_VERSION__)
int variant1838_4(int value, ...) {
    return value + 65;
}
#else
int variant1838_4(int value) {
    return value + 65;
}
#endif
struct Node1838_5 {
    int value;
    struct Node1838_5 *next;
};

int length1838_5(const struct Node1838_5 *head) {
    int count = 0;
    const struct Node1838_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
struct Node1838_6 {
    int value;
    struct Node1838_6 *next;
};

int length1838_6(const struct Node1838_6 *head) {
    int count = 0;
    const struct Node1838_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
