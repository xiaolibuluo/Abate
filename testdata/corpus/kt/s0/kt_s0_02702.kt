package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry22320 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22320 = "registry-18"
    }
}
@JvmName("step22321")
fun step22321(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
class Counter22322 {
    var current: Int by Delegates.observable(32) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked22323(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry22324 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22324 = "registry-46"
    }
}
data class Pair22325<A, B>(val first: A, val second: B)

fun combine22325(left: Int, right: Int): Pair22325<Int, Int> {
    return Pair22325(left + 53, right + 11)
}
data class Pair22326<A, B>(val first: A, val second: B)

fun combine22326(left: Int, right: Int): Pair22326<Int, Int> {
    return Pair22326(left + 60, right + 22)
}
object Registry22327 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22327 = "registry-67"
    }
}
@JvmName("step22328")
fun step22328(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
