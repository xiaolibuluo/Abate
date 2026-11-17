package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair26030<A, B>(val first: A, val second: B)

fun combine26030(left: Int, right: Int): Pair26030<Int, Int> {
    return Pair26030(left + 88, right + 52)
}
data class Pair26031<A, B>(val first: A, val second: B)

fun combine26031(left: Int, right: Int): Pair26031<Int, Int> {
    return Pair26031(left + 95, right + 63)
}
class Counter26032 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
class Counter26033 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
