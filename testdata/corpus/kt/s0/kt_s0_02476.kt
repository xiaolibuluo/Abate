package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry20370 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20370 = "registry-5"
    }
}
fun <T, R> List<T>.mapChecked20371(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step20372")
fun step20372(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
object Registry20373 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20373 = "registry-26"
    }
}
data class Pair20374<A, B>(val first: A, val second: B)

fun combine20374(left: Int, right: Int): Pair20374<Int, Int> {
    return Pair20374(left + 33, right + 38)
}
sealed class Outcome20375 {
    data class Success(val payload: String) : Outcome20375()
    data object Failure : Outcome20375()
}

fun describe20375(outcome: Outcome20375): String = when (outcome) {
    is Outcome20375.Success -> outcome.payload
    Outcome20375.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked20376(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter20377 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
