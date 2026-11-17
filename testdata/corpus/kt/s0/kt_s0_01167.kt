package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9600<A, B>(val first: A, val second: B)

fun combine9600(left: Int, right: Int): Pair9600<Int, Int> {
    return Pair9600(left + 69, right + 19)
}
sealed class Outcome9601 {
    data class Success(val payload: String) : Outcome9601()
    data object Failure : Outcome9601()
}

fun describe9601(outcome: Outcome9601): String = when (outcome) {
    is Outcome9601.Success -> outcome.payload
    Outcome9601.Failure -> "failed"
}
@JvmName("step9602")
fun step9602(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
fun <T> List<T>.secondOrNull9603(): T? = if (size > 1) this[1] else null

fun render9603(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome9604 {
    data class Success(val payload: String) : Outcome9604()
    data object Failure : Outcome9604()
}

fun describe9604(outcome: Outcome9604): String = when (outcome) {
    is Outcome9604.Success -> outcome.payload
    Outcome9604.Failure -> "failed"
}
object Registry9605 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9605 = "registry-7"
    }
}
fun <T, R> List<T>.mapChecked9606(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
