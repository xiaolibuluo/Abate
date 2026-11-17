package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17940(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step17941")
fun step17941(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
object Registry17942 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17942 = "registry-61"
    }
}
@JvmName("step17943")
fun step17943(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 68
}
fun <T, R> List<T>.mapChecked17944(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome17945 {
    data class Success(val payload: String) : Outcome17945()
    data object Failure : Outcome17945()
}

fun describe17945(outcome: Outcome17945): String = when (outcome) {
    is Outcome17945.Success -> outcome.payload
    Outcome17945.Failure -> "failed"
}
class Counter17946 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
@JvmName("step17947")
fun step17947(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
object Registry17948 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17948 = "registry-6"
    }
}
