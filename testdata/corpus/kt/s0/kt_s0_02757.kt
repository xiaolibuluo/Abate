package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked22760(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter22761 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
data class Pair22762<A, B>(val first: A, val second: B)

fun combine22762(left: Int, right: Int): Pair22762<Int, Int> {
    return Pair22762(left + 22, right + 42)
}
object Registry22763 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22763 = "registry-29"
    }
}
object Registry22764 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22764 = "registry-36"
    }
}
fun <T, R> List<T>.mapChecked22765(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair22766<A, B>(val first: A, val second: B)

fun combine22766(left: Int, right: Int): Pair22766<Int, Int> {
    return Pair22766(left + 50, right + 25)
}
fun <T, R> List<T>.mapChecked22767(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
