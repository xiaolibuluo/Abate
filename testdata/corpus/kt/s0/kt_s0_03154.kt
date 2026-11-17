package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25910 {
    data class Success(val payload: String) : Outcome25910()
    data object Failure : Outcome25910()
}

fun describe25910(outcome: Outcome25910): String = when (outcome) {
    is Outcome25910.Success -> outcome.payload
    Outcome25910.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked25911(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry25912 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25912 = "registry-43"
    }
}
