package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12740 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12740 = "registry-77"
    }
}
fun <T, R> List<T>.mapChecked12741(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12742 {
    data class Success(val payload: String) : Outcome12742()
    data object Failure : Outcome12742()
}

fun describe12742(outcome: Outcome12742): String = when (outcome) {
    is Outcome12742.Success -> outcome.payload
    Outcome12742.Failure -> "failed"
}
