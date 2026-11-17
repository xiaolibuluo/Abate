package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked18620(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18621 {
    data class Success(val payload: String) : Outcome18621()
    data object Failure : Outcome18621()
}

fun describe18621(outcome: Outcome18621): String = when (outcome) {
    is Outcome18621.Success -> outcome.payload
    Outcome18621.Failure -> "failed"
}
object Registry18622 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18622 = "registry-72"
    }
}
