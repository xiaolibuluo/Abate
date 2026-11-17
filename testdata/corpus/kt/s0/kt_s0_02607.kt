package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21490<A, B>(val first: A, val second: B)

fun combine21490(left: Int, right: Int): Pair21490<Int, Int> {
    return Pair21490(left + 6, right + 46)
}
class Counter21491 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

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
