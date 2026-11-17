package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12890<A, B>(val first: A, val second: B)

fun combine12890(left: Int, right: Int): Pair12890<Int, Int> {
    return Pair12890(left + 78, right + 65)
}
data class Pair12891<A, B>(val first: A, val second: B)

fun combine12891(left: Int, right: Int): Pair12891<Int, Int> {
    return Pair12891(left + 85, right + 15)
}
@JvmName("step12892")
fun step12892(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
class Counter12893 {
    var current: Int by Delegates.observable(99) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step12894")
fun step12894(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
data class Pair12895<A, B>(val first: A, val second: B)

fun combine12895(left: Int, right: Int): Pair12895<Int, Int> {
    return Pair12895(left + 16, right + 59)
}
