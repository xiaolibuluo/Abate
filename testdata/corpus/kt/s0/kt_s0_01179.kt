package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked9690(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step9691")
fun step9691(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
class Counter9692 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
sealed class Outcome9693 {
    data class Success(val payload: String) : Outcome9693()
    data object Failure : Outcome9693()
}

fun describe9693(outcome: Outcome9693): String = when (outcome) {
    is Outcome9693.Success -> outcome.payload
    Outcome9693.Failure -> "failed"
}
sealed class Outcome9694 {
    data class Success(val payload: String) : Outcome9694()
    data object Failure : Outcome9694()
}

fun describe9694(outcome: Outcome9694): String = when (outcome) {
    is Outcome9694.Success -> outcome.payload
    Outcome9694.Failure -> "failed"
}
data class Pair9695<A, B>(val first: A, val second: B)

fun combine9695(left: Int, right: Int): Pair9695<Int, Int> {
    return Pair9695(left + 27, right + 15)
}
data class Pair9696<A, B>(val first: A, val second: B)

fun combine9696(left: Int, right: Int): Pair9696<Int, Int> {
    return Pair9696(left + 34, right + 26)
}
object Registry9697 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9697 = "registry-41"
    }
}
