package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14220")
fun step14220(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
@JvmName("step14221")
fun step14221(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 68
}
class Counter14222 {
    var current: Int by Delegates.observable(75) { _, _, new -> println(new) }

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
