package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter750 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
@JvmName("step751")
fun step751(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
data class Pair752<A, B>(val first: A, val second: B)

fun combine752(left: Int, right: Int): Pair752<Int, Int> {
    return Pair752(left + 24, right + 68)
}
data class Pair753<A, B>(val first: A, val second: B)

fun combine753(left: Int, right: Int): Pair753<Int, Int> {
    return Pair753(left + 31, right + 18)
}
class Counter754 {
    var current: Int by Delegates.observable(38) { _, _, new -> println(new) }

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
