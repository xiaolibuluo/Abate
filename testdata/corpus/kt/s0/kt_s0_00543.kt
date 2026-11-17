package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4400<A, B>(val first: A, val second: B)

fun combine4400(left: Int, right: Int): Pair4400<Int, Int> {
    return Pair4400(left + 99, right + 39)
}
class Counter4401 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
