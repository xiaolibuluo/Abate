
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind375_0 { KIND_EMPTY375_0 = 0, KIND_FULL375_0 = 30 };

union Cell375_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind375_0(union Cell375_0 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL375_0 : KIND_EMPTY375_0;
}
struct Node375_1 {
    int value;
    struct Node375_1 *next;
};

int length375_1(const struct Node375_1 *head) {
    int count = 0;
    const struct Node375_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
enum Kind375_2 { KIND_EMPTY375_2 = 0, KIND_FULL375_2 = 44 };

union Cell375_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind375_2(union Cell375_2 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL375_2 : KIND_EMPTY375_2;
}
int scan375_3(const int *begin, const int *end) {
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
struct Node375_4 {
    int value;
    struct Node375_4 *next;
};

int length375_4(const struct Node375_4 *head) {
    int count = 0;
    const struct Node375_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
