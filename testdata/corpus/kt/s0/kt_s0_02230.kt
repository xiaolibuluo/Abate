package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18340 {
    data class Success(val payload: String) : Outcome18340()
    data object Failure : Outcome18340()
}

fun describe18340(outcome: Outcome18340): String = when (outcome) {
    is Outcome18340.Success -> outcome.payload
    Outcome18340.Failure -> "failed"
}
sealed class Outcome18341 {
    data class Success(val payload: String) : Outcome18341()
    data object Failure : Outcome18341()
}

fun describe18341(outcome: Outcome18341): String = when (outcome) {
    is Outcome18341.Success -> outcome.payload
    Outcome18341.Failure -> "failed"
}
sealed class Outcome18342 {
    data class Success(val payload: String) : Outcome18342()
    data object Failure : Outcome18342()
}

fun describe18342(outcome: Outcome18342): String = when (outcome) {
    is Outcome18342.Success -> outcome.payload
    Outcome18342.Failure -> "failed"
}
sealed class Outcome18343 {
    data class Success(val payload: String) : Outcome18343()
    data object Failure : Outcome18343()
}

fun describe18343(outcome: Outcome18343): String = when (outcome) {
    is Outcome18343.Success -> outcome.payload
    Outcome18343.Failure -> "failed"
}
object Registry18344 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18344 = "registry-13"
    }
}
fun <T, R> List<T>.mapChecked18345(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
