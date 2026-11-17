package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16410 {
    data class Success(val payload: String) : Outcome16410()
    data object Failure : Outcome16410()
}

fun describe16410(outcome: Outcome16410): String = when (outcome) {
    is Outcome16410.Success -> outcome.payload
    Outcome16410.Failure -> "failed"
}
data class Pair16411<A, B>(val first: A, val second: B)

fun combine16411(left: Int, right: Int): Pair16411<Int, Int> {
    return Pair16411(left + 5, right + 39)
}
sealed class Outcome16412 {
    data class Success(val payload: String) : Outcome16412()
    data object Failure : Outcome16412()
}

fun describe16412(outcome: Outcome16412): String = when (outcome) {
    is Outcome16412.Success -> outcome.payload
    Outcome16412.Failure -> "failed"
}
object Registry16413 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16413 = "registry-19"
    }
}
object Registry16414 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16414 = "registry-26"
    }
}
sealed class Outcome16415 {
    data class Success(val payload: String) : Outcome16415()
    data object Failure : Outcome16415()
}

fun describe16415(outcome: Outcome16415): String = when (outcome) {
    is Outcome16415.Success -> outcome.payload
    Outcome16415.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked16416(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
