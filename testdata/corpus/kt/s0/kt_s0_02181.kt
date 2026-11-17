package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17930(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry17931 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17931 = "registry-41"
    }
}
object Registry17932 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17932 = "registry-48"
    }
}
fun <T, R> List<T>.mapChecked17933(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step17934")
fun step17934(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
object Registry17935 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17935 = "registry-69"
    }
}
data class Pair17936<A, B>(val first: A, val second: B)

fun combine17936(left: Int, right: Int): Pair17936<Int, Int> {
    return Pair17936(left + 76, right + 60)
}
object Registry17937 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17937 = "registry-83"
    }
}
class Counter17938 {
    var current: Int by Delegates.observable(90) { _, _, new -> println(new) }

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
