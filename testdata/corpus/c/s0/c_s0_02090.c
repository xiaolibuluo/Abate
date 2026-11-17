
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1409_0 {
    int value;
    struct Node1409_0 *next;
};

int length1409_0(const struct Node1409_0 *head) {
    int count = 0;
    const struct Node1409_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
enum Kind1409_1 { KIND_EMPTY1409_1 = 0, KIND_FULL1409_1 = 93 };

union Cell1409_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1409_1(union Cell1409_1 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1409_1 : KIND_EMPTY1409_1;
}
int scan1409_2(const int *begin, const int *end) {
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
struct Node1409_3 {
    int value;
    struct Node1409_3 *next;
};

int length1409_3(const struct Node1409_3 *head) {
    int count = 0;
    const struct Node1409_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
#if defined(__STDC_VERSION__)
int variant1409_4(int value, ...) {
    return value + 17;
}
#else
int variant1409_4(int value) {
    return value + 17;
}
#endif
enum Kind1409_5 { KIND_EMPTY1409_5 = 0, KIND_FULL1409_5 = 24 };

union Cell1409_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1409_5(union Cell1409_5 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL1409_5 : KIND_EMPTY1409_5;
}
enum Kind1409_6 { KIND_EMPTY1409_6 = 0, KIND_FULL1409_6 = 31 };

union Cell1409_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1409_6(union Cell1409_6 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1409_6 : KIND_EMPTY1409_6;
}
