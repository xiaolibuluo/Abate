package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23090<A, B>(val first: A, val second: B)

fun combine23090(left: Int, right: Int): Pair23090<Int, Int> {
    return Pair23090(left + 49, right + 68)
}
object Registry23091 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23091 = "registry-56"
    }
}
object Registry23092 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23092 = "registry-63"
    }
}
object Registry23093 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23093 = "registry-70"
    }
}
fun <T, R> List<T>.mapChecked23094(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter23095 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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
