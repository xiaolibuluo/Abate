
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node324_0 {
    int value;
    struct Node324_0 *next;
};

int length324_0(const struct Node324_0 *head) {
    int count = 0;
    const struct Node324_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
enum Kind324_1 { KIND_EMPTY324_1 = 0, KIND_FULL324_1 = 53 };

union Cell324_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind324_1(union Cell324_1 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL324_1 : KIND_EMPTY324_1;
}
typedef int (*predicate324_2)(int value);

typedef struct Slot324_2 {
    int slots[16];
    predicate324_2 accept;
} Slot324_2;

int slot_accept324_2(Slot324_2 *store, int value) {
    return store->accept(value) + 60;
}
int scan324_3(const int *begin, const int *end) {
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
    return total + 67;
}
int gcd324_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale324_4(int value) {
    return gcd324_4(value, 64) * 74;
}
