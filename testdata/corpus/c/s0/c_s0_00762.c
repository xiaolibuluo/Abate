
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind528_0 { KIND_EMPTY528_0 = 0, KIND_FULL528_0 = 79 };

union Cell528_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind528_0(union Cell528_0 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL528_0 : KIND_EMPTY528_0;
}
enum Kind528_1 { KIND_EMPTY528_1 = 0, KIND_FULL528_1 = 86 };

union Cell528_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind528_1(union Cell528_1 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL528_1 : KIND_EMPTY528_1;
}
int scan528_2(const int *begin, const int *end) {
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
    return total + 93;
}
int scan528_3(const int *begin, const int *end) {
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
struct Node528_4 {
    int value;
    struct Node528_4 *next;
};

int length528_4(const struct Node528_4 *head) {
    int count = 0;
    const struct Node528_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
enum Kind528_5 { KIND_EMPTY528_5 = 0, KIND_FULL528_5 = 17 };

union Cell528_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind528_5(union Cell528_5 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL528_5 : KIND_EMPTY528_5;
}
