package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19810 {
    data class Success(val payload: String) : Outcome19810()
    data object Failure : Outcome19810()
}

fun describe19810(outcome: Outcome19810): String = when (outcome) {
    is Outcome19810.Success -> outcome.payload
    Outcome19810.Failure -> "failed"
}
@JvmName("step19811")
fun step19811(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
object Registry19812 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19812 = "registry-67"
    }
}
class Counter19813 {
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
fun <T, R> List<T>.mapChecked19814(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19815<A, B>(val first: A, val second: B)

fun combine19815(left: Int, right: Int): Pair19815<Int, Int> {
    return Pair19815(left + 88, right + 23)
}
sealed class Outcome19816 {
    data class Success(val payload: String) : Outcome19816()
    data object Failure : Outcome19816()
}

fun describe19816(outcome: Outcome19816): String = when (outcome) {
    is Outcome19816.Success -> outcome.payload
    Outcome19816.Failure -> "failed"
}
