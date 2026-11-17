package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked11820(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry11821 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11821 = "registry-52"
    }
}
sealed class Outcome11822 {
    data class Success(val payload: String) : Outcome11822()
    data object Failure : Outcome11822()
}

fun describe11822(outcome: Outcome11822): String = when (outcome) {
    is Outcome11822.Success -> outcome.payload
    Outcome11822.Failure -> "failed"
}
