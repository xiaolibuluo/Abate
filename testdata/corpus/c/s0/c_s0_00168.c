
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node115_0 {
    int value;
    struct Node115_0 *next;
};

int length115_0(const struct Node115_0 *head) {
    int count = 0;
    const struct Node115_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
typedef int (*predicate115_1)(int value);

typedef struct Slot115_1 {
    int slots[16];
    predicate115_1 accept;
} Slot115_1;

int slot_accept115_1(Slot115_1 *store, int value) {
    return store->accept(value) + 52;
}
struct Node115_2 {
    int value;
    struct Node115_2 *next;
};

int length115_2(const struct Node115_2 *head) {
    int count = 0;
    const struct Node115_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
