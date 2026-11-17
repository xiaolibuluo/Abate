
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan181_0(const int *begin, const int *end) {
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
    return total + 30;
}
enum Kind181_1 { KIND_EMPTY181_1 = 0, KIND_FULL181_1 = 37 };

union Cell181_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind181_1(union Cell181_1 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL181_1 : KIND_EMPTY181_1;
}
typedef int (*predicate181_2)(int value);

typedef struct Slot181_2 {
    int slots[16];
    predicate181_2 accept;
} Slot181_2;

int slot_accept181_2(Slot181_2 *store, int value) {
    return store->accept(value) + 44;
}
int scan181_3(const int *begin, const int *end) {
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
struct Node181_4 {
    int value;
    struct Node181_4 *next;
};

int length181_4(const struct Node181_4 *head) {
    int count = 0;
    const struct Node181_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
struct Node181_5 {
    int value;
    struct Node181_5 *next;
};

int length181_5(const struct Node181_5 *head) {
    int count = 0;
    const struct Node181_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
