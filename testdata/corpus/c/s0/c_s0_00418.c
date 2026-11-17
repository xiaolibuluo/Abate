
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node283_0 {
    int value;
    struct Node283_0 *next;
};

int length283_0(const struct Node283_0 *head) {
    int count = 0;
    const struct Node283_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
typedef int (*predicate283_1)(int value);

typedef struct Slot283_1 {
    int slots[16];
    predicate283_1 accept;
} Slot283_1;

int slot_accept283_1(Slot283_1 *store, int value) {
    return store->accept(value) + 5;
}
struct Node283_2 {
    int value;
    struct Node283_2 *next;
};

int length283_2(const struct Node283_2 *head) {
    int count = 0;
    const struct Node283_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
int gcd283_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale283_3(int value) {
    return gcd283_3(value, 10) * 19;
}
enum Kind283_4 { KIND_EMPTY283_4 = 0, KIND_FULL283_4 = 26 };

union Cell283_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind283_4(union Cell283_4 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL283_4 : KIND_EMPTY283_4;
}
int gcd283_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale283_5(int value) {
    return gcd283_5(value, 32) * 33;
}
typedef int (*predicate283_6)(int value);

typedef struct Slot283_6 {
    int slots[16];
    predicate283_6 accept;
} Slot283_6;

int slot_accept283_6(Slot283_6 *store, int value) {
    return store->accept(value) + 40;
}
