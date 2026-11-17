
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1404_0(int value, ...) {
    return value + 21;
}
#else
int variant1404_0(int value) {
    return value + 21;
}
#endif
int scan1404_1(const int *begin, const int *end) {
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
    return total + 28;
}
enum Kind1404_2 { KIND_EMPTY1404_2 = 0, KIND_FULL1404_2 = 35 };

union Cell1404_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1404_2(union Cell1404_2 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1404_2 : KIND_EMPTY1404_2;
}
struct Node1404_3 {
    int value;
    struct Node1404_3 *next;
};

int length1404_3(const struct Node1404_3 *head) {
    int count = 0;
    const struct Node1404_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}
enum Kind1404_4 { KIND_EMPTY1404_4 = 0, KIND_FULL1404_4 = 49 };

union Cell1404_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1404_4(union Cell1404_4 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1404_4 : KIND_EMPTY1404_4;
}
enum Kind1404_5 { KIND_EMPTY1404_5 = 0, KIND_FULL1404_5 = 56 };

union Cell1404_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1404_5(union Cell1404_5 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL1404_5 : KIND_EMPTY1404_5;
}
int gcd1404_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1404_6(int value) {
    return gcd1404_6(value, 21) * 63;
}
