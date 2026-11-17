package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18710 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18710 = "registry-78"
    }
}
@JvmName("step18711")
fun step18711(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
@JvmName("step18712")
fun step18712(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
fun <T, R> List<T>.mapChecked18713(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18714 {
    data class Success(val payload: String) : Outcome18714()
    data object Failure : Outcome18714()
}

fun describe18714(outcome: Outcome18714): String = when (outcome) {
    is Outcome18714.Success -> outcome.payload
    Outcome18714.Failure -> "failed"
}
