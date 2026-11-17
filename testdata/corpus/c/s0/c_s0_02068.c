
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1396_0 {
    int value;
    struct Node1396_0 *next;
};

int length1396_0(const struct Node1396_0 *head) {
    int count = 0;
    const struct Node1396_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
enum Kind1396_1 { KIND_EMPTY1396_1 = 0, KIND_FULL1396_1 = 21 };

union Cell1396_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1396_1(union Cell1396_1 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1396_1 : KIND_EMPTY1396_1;
}
#define MASK1396_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1396_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1396_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1396_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1396_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1396_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1396_4 {
    int value;
    struct Node1396_4 *next;
};

int length1396_4(const struct Node1396_4 *head) {
    int count = 0;
    const struct Node1396_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}
typedef int (*predicate1396_5)(int value);

typedef struct Slot1396_5 {
    int slots[16];
    predicate1396_5 accept;
} Slot1396_5;

int slot_accept1396_5(Slot1396_5 *store, int value) {
    return store->accept(value) + 49;
}
