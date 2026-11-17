package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry22660 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22660 = "registry-72"
    }
}
fun <T, R> List<T>.mapChecked22661(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step22662")
fun step22662(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
@JvmName("step22663")
fun step22663(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
fun <T, R> List<T>.mapChecked22664(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter22665 {
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
