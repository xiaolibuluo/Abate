
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1037_0 {
    int value;
    struct Node1037_0 *next;
};

int length1037_0(const struct Node1037_0 *head) {
    int count = 0;
    const struct Node1037_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
struct Node1037_1 {
    int value;
    struct Node1037_1 *next;
};

int length1037_1(const struct Node1037_1 *head) {
    int count = 0;
    const struct Node1037_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
struct Node1037_2 {
    int value;
    struct Node1037_2 *next;
};

int length1037_2(const struct Node1037_2 *head) {
    int count = 0;
    const struct Node1037_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
