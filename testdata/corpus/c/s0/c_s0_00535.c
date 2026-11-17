
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant367_0(int value, ...) {
    return value + 23;
}
#else
int variant367_0(int value) {
    return value + 23;
}
#endif
int scan367_1(const int *begin, const int *end) {
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
typedef int (*predicate367_2)(int value);

typedef struct Slot367_2 {
    int slots[16];
    predicate367_2 accept;
} Slot367_2;

int slot_accept367_2(Slot367_2 *store, int value) {
    return store->accept(value) + 37;
}
enum Kind367_3 { KIND_EMPTY367_3 = 0, KIND_FULL367_3 = 44 };

union Cell367_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind367_3(union Cell367_3 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL367_3 : KIND_EMPTY367_3;
}
struct Node367_4 {
    int value;
    struct Node367_4 *next;
};

int length367_4(const struct Node367_4 *head) {
    int count = 0;
    const struct Node367_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
