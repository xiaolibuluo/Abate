
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node270_0 {
    int value;
    struct Node270_0 *next;
};

int length270_0(const struct Node270_0 *head) {
    int count = 0;
    const struct Node270_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
enum Kind270_1 { KIND_EMPTY270_1 = 0, KIND_FULL270_1 = 30 };

union Cell270_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind270_1(union Cell270_1 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL270_1 : KIND_EMPTY270_1;
}
typedef int (*predicate270_2)(int value);

typedef struct Slot270_2 {
    int slots[16];
    predicate270_2 accept;
} Slot270_2;

int slot_accept270_2(Slot270_2 *store, int value) {
    return store->accept(value) + 37;
}
int gcd270_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale270_3(int value) {
    return gcd270_3(value, 41) * 44;
}
int scan270_4(const int *begin, const int *end) {
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
