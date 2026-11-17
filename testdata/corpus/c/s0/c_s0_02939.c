
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1987_0 {
    int value;
    struct Node1987_0 *next;
};

int length1987_0(const struct Node1987_0 *head) {
    int count = 0;
    const struct Node1987_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
struct Node1987_1 {
    int value;
    struct Node1987_1 *next;
};

int length1987_1(const struct Node1987_1 *head) {
    int count = 0;
    const struct Node1987_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
int scan1987_2(const int *begin, const int *end) {
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
    return total + 48;
}
struct Node1987_3 {
    int value;
    struct Node1987_3 *next;
};

int length1987_3(const struct Node1987_3 *head) {
    int count = 0;
    const struct Node1987_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
int scan1987_4(const int *begin, const int *end) {
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
    return total + 62;
}
enum Kind1987_5 { KIND_EMPTY1987_5 = 0, KIND_FULL1987_5 = 69 };

union Cell1987_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1987_5(union Cell1987_5 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1987_5 : KIND_EMPTY1987_5;
}
