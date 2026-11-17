package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25190<A, B>(val first: A, val second: B)

fun combine25190(left: Int, right: Int): Pair25190<Int, Int> {
    return Pair25190(left + 63, right + 13)
}
sealed class Outcome25191 {
    data class Success(val payload: String) : Outcome25191()
    data object Failure : Outcome25191()
}

fun describe25191(outcome: Outcome25191): String = when (outcome) {
    is Outcome25191.Success -> outcome.payload
    Outcome25191.Failure -> "failed"
}
sealed class Outcome25192 {
    data class Success(val payload: String) : Outcome25192()
    data object Failure : Outcome25192()
}

fun describe25192(outcome: Outcome25192): String = when (outcome) {
    is Outcome25192.Success -> outcome.payload
    Outcome25192.Failure -> "failed"
}
object Registry25193 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25193 = "registry-84"
    }
}
sealed class Outcome25194 {
    data class Success(val payload: String) : Outcome25194()
    data object Failure : Outcome25194()
}

fun describe25194(outcome: Outcome25194): String = when (outcome) {
    is Outcome25194.Success -> outcome.payload
    Outcome25194.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25195(): T? = if (size > 1) this[1] else null

fun render25195(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter25196 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
