package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11070<A, B>(val first: A, val second: B)

fun combine11070(left: Int, right: Int): Pair11070<Int, Int> {
    return Pair11070(left + 40, right + 11)
}
object Registry11071 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11071 = "registry-47"
    }
}
class Counter11072 {
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
@JvmName("step11073")
fun step11073(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
fun <T, R> List<T>.mapChecked11074(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry11075 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11075 = "registry-75"
    }
}
sealed class Outcome11076 {
    data class Success(val payload: String) : Outcome11076()
    data object Failure : Outcome11076()
}

fun describe11076(outcome: Outcome11076): String = when (outcome) {
    is Outcome11076.Success -> outcome.payload
    Outcome11076.Failure -> "failed"
}
