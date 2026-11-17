package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked24740(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair24741<A, B>(val first: A, val second: B)

fun combine24741(left: Int, right: Int): Pair24741<Int, Int> {
    return Pair24741(left + 67, right + 14)
}
class Counter24742 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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
data class Pair24743<A, B>(val first: A, val second: B)

fun combine24743(left: Int, right: Int): Pair24743<Int, Int> {
    return Pair24743(left + 81, right + 36)
}
object Registry24744 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24744 = "registry-88"
    }
}
object Registry24745 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24745 = "registry-95"
    }
}
fun <T, R> List<T>.mapChecked24746(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry24747 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24747 = "registry-12"
    }
}
object Registry24748 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24748 = "registry-19"
    }
}
