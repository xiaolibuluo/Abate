package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15660")
fun step15660(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
class Counter15661 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
