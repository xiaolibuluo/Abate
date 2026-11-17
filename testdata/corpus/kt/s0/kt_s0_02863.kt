package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23620 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23620 = "registry-59"
    }
}
fun <T, R> List<T>.mapChecked23621(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair23622<A, B>(val first: A, val second: B)

fun combine23622(left: Int, right: Int): Pair23622<Int, Int> {
    return Pair23622(left + 73, right + 34)
}
object Registry23623 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23623 = "registry-80"
    }
}
data class Pair23624<A, B>(val first: A, val second: B)

fun combine23624(left: Int, right: Int): Pair23624<Int, Int> {
    return Pair23624(left + 87, right + 56)
}
class Counter23625 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

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
