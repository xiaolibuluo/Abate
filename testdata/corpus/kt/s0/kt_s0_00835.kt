package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6860")
fun step6860(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
class Counter6861 {
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
