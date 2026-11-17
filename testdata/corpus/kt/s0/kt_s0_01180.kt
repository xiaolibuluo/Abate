package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9700<A, B>(val first: A, val second: B)

fun combine9700(left: Int, right: Int): Pair9700<Int, Int> {
    return Pair9700(left + 5, right + 28)
}
sealed class Outcome9701 {
    data class Success(val payload: String) : Outcome9701()
    data object Failure : Outcome9701()
}

fun describe9701(outcome: Outcome9701): String = when (outcome) {
    is Outcome9701.Success -> outcome.payload
    Outcome9701.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked9702(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry9703 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9703 = "registry-26"
    }
}
data class Pair9704<A, B>(val first: A, val second: B)

fun combine9704(left: Int, right: Int): Pair9704<Int, Int> {
    return Pair9704(left + 33, right + 11)
}
fun <T> List<T>.secondOrNull9705(): T? = if (size > 1) this[1] else null

fun render9705(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9706")
fun step9706(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
sealed class Outcome9707 {
    data class Success(val payload: String) : Outcome9707()
    data object Failure : Outcome9707()
}

fun describe9707(outcome: Outcome9707): String = when (outcome) {
    is Outcome9707.Success -> outcome.payload
    Outcome9707.Failure -> "failed"
}
class Counter9708 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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
