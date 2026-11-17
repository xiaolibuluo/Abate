package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14620<A, B>(val first: A, val second: B)

fun combine14620(left: Int, right: Int): Pair14620<Int, Int> {
    return Pair14620(left + 96, right + 56)
}
class Counter14621 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
