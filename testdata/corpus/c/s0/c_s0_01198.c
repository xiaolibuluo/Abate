
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node821_0 {
    int value;
    struct Node821_0 *next;
};

int length821_0(const struct Node821_0 *head) {
    int count = 0;
    const struct Node821_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
enum Kind821_1 { KIND_EMPTY821_1 = 0, KIND_FULL821_1 = 15 };

union Cell821_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind821_1(union Cell821_1 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL821_1 : KIND_EMPTY821_1;
}
enum Kind821_2 { KIND_EMPTY821_2 = 0, KIND_FULL821_2 = 22 };

union Cell821_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind821_2(union Cell821_2 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL821_2 : KIND_EMPTY821_2;
}
struct Node821_3 {
    int value;
    struct Node821_3 *next;
};

int length821_3(const struct Node821_3 *head) {
    int count = 0;
    const struct Node821_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
typedef int (*predicate821_4)(int value);

typedef struct Slot821_4 {
    int slots[16];
    predicate821_4 accept;
} Slot821_4;

int slot_accept821_4(Slot821_4 *store, int value) {
    return store->accept(value) + 36;
}
int scan821_5(const int *begin, const int *end) {
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
    return total + 43;
}
int scan821_6(const int *begin, const int *end) {
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
    return total + 50;
}
