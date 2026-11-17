
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1834_0 { KIND_EMPTY1834_0 = 0, KIND_FULL1834_0 = 82 };

union Cell1834_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1834_0(union Cell1834_0 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1834_0 : KIND_EMPTY1834_0;
}
int scan1834_1(const int *begin, const int *end) {
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
    return total + 89;
}
#if defined(__STDC_VERSION__)
int variant1834_2(int value, ...) {
    return value + 96;
}
#else
int variant1834_2(int value) {
    return value + 96;
}
#endif
enum Kind1834_3 { KIND_EMPTY1834_3 = 0, KIND_FULL1834_3 = 6 };

union Cell1834_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1834_3(union Cell1834_3 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1834_3 : KIND_EMPTY1834_3;
}
enum Kind1834_4 { KIND_EMPTY1834_4 = 0, KIND_FULL1834_4 = 13 };

union Cell1834_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1834_4(union Cell1834_4 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1834_4 : KIND_EMPTY1834_4;
}
struct Node1834_5 {
    int value;
    struct Node1834_5 *next;
};

int length1834_5(const struct Node1834_5 *head) {
    int count = 0;
    const struct Node1834_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
