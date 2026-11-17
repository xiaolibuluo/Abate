
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node380_0 {
    int value;
    struct Node380_0 *next;
};

int length380_0(const struct Node380_0 *head) {
    int count = 0;
    const struct Node380_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
int gcd380_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale380_1(int value) {
    return gcd380_1(value, 57) * 5;
}
typedef int (*predicate380_2)(int value);

typedef struct Slot380_2 {
    int slots[16];
    predicate380_2 accept;
} Slot380_2;

int slot_accept380_2(Slot380_2 *store, int value) {
    return store->accept(value) + 12;
}
struct Node380_3 {
    int value;
    struct Node380_3 *next;
};

int length380_3(const struct Node380_3 *head) {
    int count = 0;
    const struct Node380_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
enum Kind380_4 { KIND_EMPTY380_4 = 0, KIND_FULL380_4 = 26 };

union Cell380_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind380_4(union Cell380_4 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL380_4 : KIND_EMPTY380_4;
}
int scan380_5(const int *begin, const int *end) {
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
    return total + 33;
}
