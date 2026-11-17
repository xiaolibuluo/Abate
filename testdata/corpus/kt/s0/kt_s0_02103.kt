package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter17360 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked17361(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry17362 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17362 = "registry-83"
    }
}
object Registry17363 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17363 = "registry-90"
    }
}
@JvmName("step17364")
fun step17364(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
fun <T, R> List<T>.mapChecked17365(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step17366")
fun step17366(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
