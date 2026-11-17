
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1118_0(const int *begin, const int *end) {
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
    return total + 86;
}
struct Node1118_1 {
    int value;
    struct Node1118_1 *next;
};

int length1118_1(const struct Node1118_1 *head) {
    int count = 0;
    const struct Node1118_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
int scan1118_2(const int *begin, const int *end) {
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
    return total + 100;
}
int scan1118_3(const int *begin, const int *end) {
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
    return total + 10;
}
enum Kind1118_4 { KIND_EMPTY1118_4 = 0, KIND_FULL1118_4 = 17 };

union Cell1118_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1118_4(union Cell1118_4 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL1118_4 : KIND_EMPTY1118_4;
}
enum Kind1118_5 { KIND_EMPTY1118_5 = 0, KIND_FULL1118_5 = 24 };

union Cell1118_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1118_5(union Cell1118_5 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL1118_5 : KIND_EMPTY1118_5;
}
int scan1118_6(const int *begin, const int *end) {
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
    return total + 31;
}
