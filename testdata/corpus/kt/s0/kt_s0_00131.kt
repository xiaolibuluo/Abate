package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked1060(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry1061 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1061 = "registry-32"
    }
}
sealed class Outcome1062 {
    data class Success(val payload: String) : Outcome1062()
    data object Failure : Outcome1062()
}

fun describe1062(outcome: Outcome1062): String = when (outcome) {
    is Outcome1062.Success -> outcome.payload
    Outcome1062.Failure -> "failed"
}
