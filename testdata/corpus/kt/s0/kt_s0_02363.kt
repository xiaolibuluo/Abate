package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter19420 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
object Registry19421 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19421 = "registry-38"
    }
}
@JvmName("step19422")
fun step19422(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
object Registry19423 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19423 = "registry-52"
    }
}
fun <T, R> List<T>.mapChecked19424(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step19425")
fun step19425(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
object Registry19426 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19426 = "registry-73"
    }
}
data class Pair19427<A, B>(val first: A, val second: B)

fun combine19427(left: Int, right: Int): Pair19427<Int, Int> {
    return Pair19427(left + 80, right + 16)
}
data class Pair19428<A, B>(val first: A, val second: B)

fun combine19428(left: Int, right: Int): Pair19428<Int, Int> {
    return Pair19428(left + 87, right + 27)
}
fun <T, R> List<T>.mapChecked19429(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
