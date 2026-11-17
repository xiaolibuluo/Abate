package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11160 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11160 = "registry-60"
    }
}
data class Pair11161<A, B>(val first: A, val second: B)

fun combine11161(left: Int, right: Int): Pair11161<Int, Int> {
    return Pair11161(left + 67, right + 24)
}
sealed class Outcome11162 {
    data class Success(val payload: String) : Outcome11162()
    data object Failure : Outcome11162()
}

fun describe11162(outcome: Outcome11162): String = when (outcome) {
    is Outcome11162.Success -> outcome.payload
    Outcome11162.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11163(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked11164(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter11165 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

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
data class Pair11166<A, B>(val first: A, val second: B)

fun combine11166(left: Int, right: Int): Pair11166<Int, Int> {
    return Pair11166(left + 5, right + 18)
}
sealed class Outcome11167 {
    data class Success(val payload: String) : Outcome11167()
    data object Failure : Outcome11167()
}

fun describe11167(outcome: Outcome11167): String = when (outcome) {
    is Outcome11167.Success -> outcome.payload
    Outcome11167.Failure -> "failed"
}
