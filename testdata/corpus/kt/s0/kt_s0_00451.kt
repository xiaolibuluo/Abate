package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3680 {
    data class Success(val payload: String) : Outcome3680()
    data object Failure : Outcome3680()
}

fun describe3680(outcome: Outcome3680): String = when (outcome) {
    is Outcome3680.Success -> outcome.payload
    Outcome3680.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3681(): T? = if (size > 1) this[1] else null

fun render3681(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair3682<A, B>(val first: A, val second: B)

fun combine3682(left: Int, right: Int): Pair3682<Int, Int> {
    return Pair3682(left + 50, right + 45)
}
sealed class Outcome3683 {
    data class Success(val payload: String) : Outcome3683()
    data object Failure : Outcome3683()
}

fun describe3683(outcome: Outcome3683): String = when (outcome) {
    is Outcome3683.Success -> outcome.payload
    Outcome3683.Failure -> "failed"
}
@JvmName("step3684")
fun step3684(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
data class Pair3685<A, B>(val first: A, val second: B)

fun combine3685(left: Int, right: Int): Pair3685<Int, Int> {
    return Pair3685(left + 71, right + 17)
}
